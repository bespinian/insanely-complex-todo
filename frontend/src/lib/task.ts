
export default class Task {
    id: string = crypto.randomUUID()
    name: string = ''
    complete: boolean = false

    public constructor(init?: Partial<Task>) {
        Object.assign(this, init);
    }
}
