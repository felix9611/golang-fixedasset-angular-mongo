export interface AssetListDto {
    assetCode?: string
    assetName?: string
    typeIds?: string[]
    placeIds?: string[]
    deptIds?: string[]
    purchaseDates?: string[]
    page: number
    limit: number
}