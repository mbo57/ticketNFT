package crud

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "net/http"
    "regexp"

    "app/typefile"
    "app/utility"

    "golang.org/x/exp/slices"
)

type Response struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

// レスポンスを書く変数
func ResponseWriter(w http.ResponseWriter, code int, mess string) {
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
    re, err := regexp.Compile(`/(?P<table>.*)/(?P<tmp>\D*)(?P<id>[0-9]*)$`)
    if err != nil {
        ResponseWriter(w, 424, err.Error())
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
    table := matches[0][re.SubexpIndex("table")]

    fmt.Printf("tmp : %s, id : %s, table : %s\n", tmp, id, table)
    
    if tmp == "" && id == "" {
        switch r.Method {
        // tmp が空文字で id も空文字 かつ MethodがGET    -> ReadAll
        case "GET":
            ReadAll(w, r, table)
        // tmp が空文字で id も空文字 かつ MethodがPOST   -> Create
        case "POST":
            Create(w, r, table)
        default:
            ResponseWriter(w, 422, fmt.Sprintf("Methods error : un used '%s' method", r.Method))
        }
    } else if tmp == "" && id != "" {
        switch r.Method {
        // tmp が空文字で id がある　 かつ MethodがGET    -> Read
        case "GET":
            Read(w, r, id, table)
        // tmp が空文字で id がある　 かつ MethodがPATCH  -> Update
        case "PATCH":
            Update(w, r, id, table)
        // tmp が空文字で id がある　 かつ MethodがDELETE -> Delete
        case "DELETE":
            Delete(w, r, id, table)
        default:
            ResponseWriter(w, 422, fmt.Sprintf("Methods error : un used '%s' method", r.Method))
        }
    } else {
        http.NotFound(w, r)
        return
    }
    
}

// クエリレスポンスからstaffデータの抽出
func ReadDateStaff(rows *sql.Rows) ([]interface{}, error){
    var staffs []interface{}
    for rows.Next() {
        staff := typefile.Staff{}
        err := rows.Scan(&staff.Id,
                        &staff.Email,
                        &staff.Password,
                        &staff.Name)
        if err != nil {
            return nil, err
        }
        staffs = append(staffs, staff)
    }
    return staffs, nil
}

// クエリレスポンスからeventデータの抽出
func ReadDateEvent(rows *sql.Rows) ([]interface{}, error){
    var events []interface{}
    for rows.Next() {
        event := typefile.Event{}
        err := rows.Scan(&event.Id,
                        &event.Name,
                        &event.Img,
                        &event.Date,
                        &event.Venue,
                        &event.Castid,
                        &event.EventCategoryId,
                        &event.Description)
        if err != nil {
            return nil, err
        }
        events = append(events, event)
    }
    return events, nil
}

// クエリレスポンスからeventcategoryデータの抽出
func ReadDateEventCategory(rows *sql.Rows) ([]interface{}, error){
    var eventcategorys []interface{}
    for rows.Next() {
        event := typefile.EventCategory{}
        err := rows.Scan(&event.Id,
                        &event.Name)
        if err != nil {
            return nil, err
        }
        eventcategorys = append(eventcategorys, event)
    }
    return eventcategorys, nil
}

// クエリレスポンスからcastデータの抽出
func ReadDateEventCast(rows *sql.Rows) ([]interface{}, error){
    var Casts []interface{}
    for rows.Next() {
        event := typefile.Cast{}
        err := rows.Scan(&event.Id,
                        &event.Name)
        if err != nil {
            return nil, err
        }
        Casts = append(Casts, event)
    }
    return Casts, nil
}

// すべてのレコードを取得する
func ReadAll(w http.ResponseWriter, r *http.Request, table string) {
    query := fmt.Sprintf("select * from %s", table)
    
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
        if !slices.Contains(typefile.Columns[table], k) {
            ResponseWriter(w, 422, fmt.Sprintf("'%s' column not found", k))
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
        ResponseWriter(w, 423,  err.Error())
        return
    }

    // クエリレスポンスからデータを抽出し、ans に渡す
    var ans []interface{}
    switch table {
        case "staff":
            ans, err = ReadDateStaff(rows)
        case "event":
            ans, err = ReadDateEvent(rows)
        case "eventcategory":
            ans, err = ReadDateEventCategory(rows)
        case "cast":
            ans, err = ReadDateEventCast(rows)
    }
    
    if err != nil {
        ResponseWriter(w, 423, err.Error())
        return
    }

    // クエリレスポンスを閉じる
    rows.Close()
    
    // httpレスポンスを設定する
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    json.NewEncoder(w).Encode(ans)
}

// idからレコードを検索する
func Read(w http.ResponseWriter, r *http.Request, id string, table string) {
    // var staffs typefile.Staffs
    // クエリ文実行
    query := fmt.Sprintf("select * from %s where id = ?", table)
    rows, err := utility.Db.Query(query, id)
    if err != nil {
        ResponseWriter(w, 423, err.Error())
        return
    }

    // クエリレスポンスからデータを抽出し、ans に渡す
    var ans []interface{}
    switch table {
        case "staff":
            ans, err = ReadDateStaff(rows)
        case "event":
            ans, err = ReadDateEvent(rows)
        case "eventcategory":
            ans, err = ReadDateEventCategory(rows)
        case "cast":
            ans, err = ReadDateEventCast(rows)
    }
    
    if err != nil {
        ResponseWriter(w, 423, err.Error())
        return
    }

    // クエリレスポンスを閉じる
    rows.Close()
    
    // httpレスポンスを設定する
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    json.NewEncoder(w).Encode(ans)
}


func Create(w http.ResponseWriter, r *http.Request, table string) {
    // レコードを作るクエリを実行
    var result sql.Result
    var err error
    switch table {
        case "staff":
            // フォームの内容を取得
            email := r.FormValue("email")
            password := r.FormValue("password")
            name := r.FormValue("name")
            // クエリを実行
            result, err = utility.Db.Exec("INSERT INTO staff(email, password, name) VALUES(?, ?, ?)", email, password, name)
        case "event":
            // フォームの内容を取得
            name := r.FormValue("name")
            img := r.FormValue("img")
            date := r.FormValue("date")
            venue := r.FormValue("venue")
            castid := r.FormValue("castid")
            eventcategoryid := r.FormValue("eventcategoryid")
            description := r.FormValue("description")
            // レコードを作るクエリを実行
            result, err = utility.Db.Exec(`INSERT INTO event (
                                            name,
                                            img,
                                            date,
                                            venue,
                                            castid,
                                            eventcategoryid,
                                            description)
                                            VALUES (?, ?, ?, ?, ?, ?, ?)`,
                                            name,
                                            img,
                                            date,
                                            venue,
                                            castid,
                                            eventcategoryid,
                                            description)
        case "eventcategory":
            // フォームの内容を取得
            name := r.FormValue("name")
            // レコードを作るクエリを実行
            result, err = utility.Db.Exec(`INSERT INTO eventcategory (
                                            name)
                                            VALUES (?)`,
                                            name)
        case "cast":
            // フォームの内容を取得
            name := r.FormValue("name")
            // レコードを作るクエリを実行
            result, err = utility.Db.Exec(`INSERT INTO cast (
                                            name)
                                            VALUES (?)`,
                                            name)
    }
    
    if err != nil {
        ResponseWriter(w, 423,  err.Error())
        return
    }
    
    // 新しく作ったレコードのidを取得
    lastId, err := result.LastInsertId()
    
    if err != nil {
        ResponseWriter(w, 423,  err.Error())
        return
    }
    
    // httpレスポンスを設定する
    ResponseWriter(w, 200, fmt.Sprintf("sucsess : id = %d", lastId))
}

// id で既ににあるレコードを編集
func Update(w http.ResponseWriter, r *http.Request, id string, table string) {
    query := fmt.Sprintf("update %s set", table)
    
    // クエリを作成
    length := len(r.Form)
    var list []any
    i := 0
    for k, v := range r.Form {
        // 存在するカラムか確認する
        if !slices.Contains(typefile.Columns[table], k) {
            ResponseWriter(w, 422, fmt.Sprintf("'%s' column not found", k))
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
    // クエリを実行する
    result, err := utility.Db.Exec(query, list...)
    if err != nil {
        ResponseWriter(w, 423, err.Error())
        return
    }

    // 変更されたレコードの種類を数える
    delete_num, err := result.RowsAffected()
    if err != nil {
        ResponseWriter(w, 423, err.Error())
        return
    }
    // 変更されたレコードが 1 なら成功でhttpレスポンスを設定する
    // それ以外はエラーでhttpレスポンスを設定する
    if delete_num == 1 {
        ResponseWriter(w, 200, "success update")
    } else if delete_num == 0 {
        ResponseWriter(w, 422, "fail update")
    }
}

// id でレコードを消去する
func Delete(w http.ResponseWriter, r *http.Request, id string, table string) {
    // クエリ文を作成
    query := fmt.Sprintf("delete from %s where id = ?", table)
    // クエリを実行
    result, err := utility.Db.Exec(query, id)
    if err != nil {
        ResponseWriter(w, 423, err.Error())
        return
    }
    // 影響のあったレコード数を取得
    delete_num, err := result.RowsAffected()
    if err != nil {
        ResponseWriter(w, 424, err.Error())
        return
    }
    // 変更されたレコードが 1 なら成功でhttpレスポンスを設定する
    // それ以外はエラーでhttpレスポンスを設定する
    if delete_num == 1 {
        ResponseWriter(w, 200, "success")
    } else if delete_num == 0 {
        ResponseWriter(w, 422, fmt.Sprintf("Not find id = %s", id))
    }
}
