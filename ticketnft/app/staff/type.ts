// イベント
export type Events = {
    name: "イベント",
    data: Event[]
}

export type Event = {
    id: number,
    name: string,
    img: string,
    date: Date,
    venue: string,
    cast: Cast,
    eventCategory: EventCategory,
    description: string,
}

// イベントカテゴリ
export type EventCategories = {
    name: "イベントカテゴリ",
    data: EventCategory[]
}

export type EventCategory = {
    id: number,
    name: string,
}

// 出演者
export type Casts = {
    name: "出演者",
    data: Cast[]
}

export type Cast = {
    id: number,
    name: string,
}

// サンプル値

export let SampleDataEventCategory: EventCategory = {
    id: 1,
    name: "音楽ライブ"
}

export let SampleDataCast: Cast = {
    id: 1,
    name: "SUPER BEAVER"
}

export let SampleDataEvent: Event = {
    id: 1,
    name: "SUPER BEAVER 都会のラクダ TOUR 2023-2024 〜 駱駝革命21 〜",
    img: "https://material.onlineticket.jp/s/image/025431/0001/000/0254310001_3.jpg",
    date: new Date(),
    venue: "広島文化学園HBGホール",
    cast: SampleDataCast,
    eventCategory: SampleDataEventCategory,
    description: "SUPER BEAVERのライブ。SUPER BEAVERの個人的に一番好きな曲は、証明です。",
}



export let SampleDataEvents = {
    name: "イベント",
    data: [SampleDataEvent, SampleDataEvent]
}