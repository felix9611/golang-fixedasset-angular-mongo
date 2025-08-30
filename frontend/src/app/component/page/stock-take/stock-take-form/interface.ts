export interface StockTakeItemDto {
    _id?: string
    stockTakeId: string
    assetId: string
    assetCode: string
    placeId: string
    status: string
    checkTime: string
    remark?: string
}

export interface StockTakeItemFromDto {
    stockTakeId: string
    assetId: string
    assetCode: string
    assetName: string
    placeId: string
    status: string
    remark?: string
}

export interface StockTakeFormEdit {
    actionName: string
    actionPlaceId: string
    remark?: string
    stockTakeItems: any
    createdTime: string
    finishTime?: string
    _id: string
    createBy: string
    status: number
}