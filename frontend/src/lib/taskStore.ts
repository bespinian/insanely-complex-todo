import Task from "./task";
import { writable } from "svelte/store";

const URL = "http://localhost:8001/api/tasks"

const tasks = writable<Task[]>([])

function createStore() {
    return {
        tasks,

        async fetch(): Promise<void> {
            return fetch(URL)
                .then((response: Response) => response.json())
                .then((json: object[]) => tasks.set(json.map((data) => new Task(data))))
        },

        async removeById(taskId: string) {
            tasks.update((tasks: Task[]) => tasks.filter((item) => item.id !== taskId))
        },

        async insert(name: string) {
            const task = new Task({ id: crypto.randomUUID(), name: name })
            tasks.update((tasks: Task[]) => [...tasks, task])
        },

        async toggle(id: string, isComplete: boolean) {
            tasks.update((tasks: Task[]) => {
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
