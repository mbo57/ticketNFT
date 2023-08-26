package event

import (
	"app/utility"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"golang.org/x/exp/slices"
)

type Event struct{
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Img              string    `json:"img"`
	Date             time.Time `json:"date"`
	Venue            string    `json:"venue"`
	Castid           int       `json:"castid"`
	EventCategoryId  int       `json:"eventcategoryid"`
	Description      string    `json:"description"`
}

var columns = []string{"id", "name", "img", "date", "venue", "castid", "eventcategoryid", "description"}

type Events []Event

type Response struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

// レスポンスを書く変数
func ResponseWriter(w http.ResponseWriter, r *http.Request, code int, mess string) {
    var res Response
    res.Code = code
    res.Message = mess
    json.NewEncoder(w).Encode(res)
}

// URLをたたかれたときにmethodを判定
// URLからidを読み取る
func Entry(w http.ResponseWriter, r *http.Request) {
    // Parameter を取得する
    err := r.ParseForm()
    if err != nil {
        var res Response
        res.Code = 423
        res.Message = err.Error()
	    json.NewEncoder(w).Encode(res)
        return
    }
    
    // 正規表現で一番最後の数値(id)と数値の前の文字列(tmp)を読み取る
    // tmp が空文字で id も空文字 かつ MethodがGET    -> ReadAll
    // tmp が空文字で id がある　 かつ MethodがGET    -> Read
    // tmp が空文字で id も空文字 かつ MethodがPOST   -> Create
    // tmp が空文字で id がある　 かつ MethodがPATCH  -> Update
    // tmp が空文字で id がある　 かつ MethodがDELETE -> Delete
    // その他                                         -> NotFound

    // 正規表現コンパイル
    re, err := regexp.Compile(`/event/(?P<tmp>\D*)(?P<id>[0-9]*)$`)
    if err != nil {
        ResponseWriter(w, r, 424, err.Error())
        return
    }
    
    // 正規表現マッチ
    matches := re.FindAllStringSubmatch(r.URL.Path, -1)

    // 正規表現が見つからなかったらNotFound
    if matches == nil {
        http.NotFound(w, r)
        return
    }

    tmp := matches[0][re.SubexpIndex("tmp")]
    id := matches[0][re.SubexpIndex("id")]
    
    if tmp == "" && id == "" {
        switch r.Method {
        // tmp が空文字で id も空文字 かつ MethodがGET    -> ReadAll
        case "GET":
            ReadAll(w, r)
        // tmp が空文字で id も空文字 かつ MethodがPOST   -> Create
        case "POST":
            Create(w, r)
        default:
            ResponseWriter(w, r, 422, fmt.Sprintf("Methods error : un used '%s' method", r.Method))
        }
    } else if tmp == "" && id != "" {
        switch r.Method {
        // tmp が空文字で id がある　 かつ MethodがGET    -> Read
        case "GET":
            Read(w, r, id)
        // tmp が空文字で id がある　 かつ MethodがPATCH  -> Update
        case "PATCH":
            Update(w, r, id)
        // tmp が空文字で id がある　 かつ MethodがDELETE -> Delete
        case "DELETE":
            Delete(w, r, id)
        default:
            ResponseWriter(w, r, 422, fmt.Sprintf("Methods error : un used '%s' method", r.Method))
        }
    } else {
        http.NotFound(w, r)
        return
    }
    
}

func ReadAll(w http.ResponseWriter, r *http.Request) {
	var events Events
    query := "select * from event"
    
    // パラメータをもとにクエリ文を作成
    length := len(r.Form)
    var list []any
    i := 0
    for k, v := range r.Form {
        // 1回目ならクエリ文にwhere句を追加
        if i == 0{
            query += " where"
        }
        // 存在するカラムか確認する
        if !slices.Contains(columns, k) {
            ResponseWriter(w, r, 422, fmt.Sprintf("'%s' column not found", k))
            return
        }
        // パラメータの値をlistに追加
        list = append(list, v[0])
        // クエリ文に" {パラメータ} = ?"を追加
        query += fmt.Sprintf(" %s = ?", k)
        // 最後以外はクエリ文に" and"を追加
        if i != length - 1 {
            query += " and"
        }
        i += 1
    }
    query += ";"
    // クエリを実行
    rows, err := utility.Db.Query(query, list...)
	if err != nil {
        ResponseWriter(w, r, 423,  err.Error())
        return
	}

    // クエリレスポンスからデータを抽出しeventsにまとめる
	for rows.Next() {
        event := Event{}
        err = rows.Scan(&event.Id,
                        &event.Name,
                        &event.Img,
                        &event.Date,
                        &event.Venue,
                        &event.Castid,
                        &event.EventCategoryId,
                        &event.Description)
		if err != nil {
            ResponseWriter(w, r, 423,  err.Error())
            return
		}
		events = append(events, event)
	}

    // クエリレスポンスを閉じる
	rows.Close()
    
    // httpレスポンスを設定する
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
	json.NewEncoder(w).Encode(events)
}

// idからレコードを検索する
func Read(w http.ResponseWriter, r *http.Request, id string) {
	var events Events
    // クエリ文実行
    rows, err := utility.Db.Query("select * from event where id = ?", id)
	if err != nil {
        ResponseWriter(w, r, 423, err.Error())
        return
	}

    // クエリレスポンスからデータの抽出
	for rows.Next() {
        event := Event{}
        err = rows.Scan(&event.Id,
                        &event.Name,
                        &event.Img,
                        &event.Date,
                        &event.Venue,
                        &event.Castid,
                        &event.EventCategoryId,
                        &event.Description)
		if err != nil {
            ResponseWriter(w, r, 423, err.Error())
            return
		}
		events = append(events, event)
	}

    // クエリレスポンスを閉じる
	rows.Close()
    
    // httpレスポンスを設定する
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
	json.NewEncoder(w).Encode(events)
}


func Create(w http.ResponseWriter, r *http.Request) {
    // フォームの内容を取得
    name := r.FormValue("name")
    img := r.FormValue("img")
    date := r.FormValue("date")
    veune := r.FormValue("veune")
    castid := r.FormValue("castid")
    eventcategoryid := r.FormValue("eventcategoryid")
    description := r.FormValue("description")

    // レコードを作るクエリを実行
    result, err := utility.Db.Exec(`INSERT INTO event (
                                    name,
                                    img,
                                    date,
                                    veune,
                                    castid,
                                    eventcategoryid,
                                    description)
                                    VALUES (?, ?, ?, ?, ?, ?, ?)`,
                                    name,
                                    img,
                                    date,
                                    veune,
                                    castid,
                                    eventcategoryid,
                                    description)
    
    if err != nil {
        ResponseWriter(w, r, 423,  err.Error())
        return
    }
    
    // 新しく作ったレコードのidを取得
    lastId, err := result.LastInsertId()
    
    if err != nil {
        ResponseWriter(w, r, 423,  err.Error())
        return
    }
    
    // httpレスポンスを設定する
    ResponseWriter(w, r, 200, fmt.Sprintf("sucsess : id = %d", lastId))
}

// id で既ににあるレコードを編集
func Update(w http.ResponseWriter, r *http.Request, id string) {
    query := "update event set"
    
    // クエリを作成
    length := len(r.Form)
    var list []any
    i := 0
    for k, v := range r.Form {
        // 存在するカラムか確認する
        if !slices.Contains(columns, k) {
            ResponseWriter(w, r, 422, fmt.Sprintf("'%s' column not found", k))
            return
        }
        // パラメータの値をlistに追加
        list = append(list, v[0])
        // クエリ文に" {パラメータ} = ?"を追加
        query += fmt.Sprintf(" %s = ?", k)
        // 最後以外はクエリ文に" ,"を追加
        if i != length - 1 {
            query += " ,"
        }
        i += 1
    }

    // id をlist に追加
    list = append(list, id)

    // クエリ文に id のwhere句を追加する
    query += " where id = ?;"
    // クエリｗを実行する
    result, err := utility.Db.Exec(query, list...)
	if err != nil {
        ResponseWriter(w, r, 423, err.Error())
        return
	}

    // 変更されたレコードの種類を数える
    delete_num, err := result.RowsAffected()
    if err != nil {
        ResponseWriter(w, r, 423, err.Error())
        return
    }
    // 変更されたレコードが 1 なら成功でhttpレスポンスを設定する
    // それ以外はエラーでhttpレスポンスを設定する
    if delete_num == 1 {
        ResponseWriter(w, r, 200, "success update")
    } else if delete_num == 0 {
        ResponseWriter(w, r, 422, "fail update")
    }
}

// id でレコードを消去する
func Delete(w http.ResponseWriter, r *http.Request, id string) {
    // クエリを実行
    result, err := utility.Db.Exec("delete from event where id = ?", id)
    if err != nil {
        ResponseWriter(w, r, 423, err.Error())
        return
    }
    // 影響のあったレコード数を取得
    delete_num, err := result.RowsAffected()
    if err != nil {
        ResponseWriter(w, r, 424, err.Error())
        return
    }
    // 変更されたレコードが 1 なら成功でhttpレスポンスを設定する
    // それ以外はエラーでhttpレスポンスを設定する
    if delete_num == 1 {
        ResponseWriter(w, r, 200, "success")
    } else if delete_num == 0 {
        ResponseWriter(w, r, 422, fmt.Sprintf("Not find id = %s", id))
    }
}
