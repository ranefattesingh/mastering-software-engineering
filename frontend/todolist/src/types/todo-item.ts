
export interface TodoItem {
    id: number | 0
    title: string | ''
    description: string | ''
    dueDate: string | ''
    isCompleted: boolean | false
    priority: string | 'Low'
    dueTime: string | ''
    createDate: string | ''
}