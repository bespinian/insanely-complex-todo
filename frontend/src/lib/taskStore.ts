import Task from "./task";
import { writable } from "svelte/store";

const demoTasks = [
    new Task({
        id: crypto.randomUUID(),
        name: 'Complete project review',
        complete: false
    }),
    new Task({
        id: crypto.randomUUID(),
        name: 'Do another task',
        complete: true
    })
]


function createStore() {
    const initialTasks: Task[] = []
    const { subscribe, update, set } = writable(initialTasks)

    return {
        subscribe,
        update,

        async init() {
            set(demoTasks)
        },

        async removeById(taskId: string) {
            update((tasks: Task[]) => tasks.filter((item) => item.id !== taskId))
        },

        async add(name: string) {
            const task = new Task({ id: crypto.randomUUID(), name: name })
            update((tasks: Task[]) => [...tasks, task])
        },

        async toggle(id: string, isComplete: boolean) {
            update((tasks: Task[]) => {
                const taskIndex = tasks.findIndex((task: Task) => task.id == id)
                if (taskIndex >= 0) {
                    tasks[taskIndex].complete = isComplete
                }
                return tasks
            })
        }
    }
}

export default createStore()
