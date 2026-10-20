export interface UserForm {
    id?: string,
    username: string
    password?: string
    avatarBase64?: string
    deptId?: number
    email: string
    roles?: string[],
    superAdmin?: boolean
    placeIds?: string[]
}