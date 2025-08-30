export interface UpdateRepairRecordDto {
    _id: string
    assetId: string
    repairReason: string
    maintenanceReriod: boolean
    maintenanceName: string
    maintenanceDate: string
    maintenanceFinishDate: string
    repairInvoiceDate: string
    repairInvoiceNo: string
    repairAmount: number
    remark: string
}