import type Task from "./task";
import { writable } from "svelte/store";

const URL = "/api/tasks"

const tasks = writable<Task[]>([])
const error = writable<string | undefined>(undefined)

function getJsonRequest(url: string, method?: string, body?: object): Request {
    return new Request(url, {
        method: method,
        body: body ? JSON.stringify(body) : null,
        headers: {
            "Content-Type": "application/json",
        },
    })
}

function createStore() {
    return {
        tasks,
        error,

        async fetch(): Promise<void> {
            return fetch(getJsonRequest(URL))
                .then((response: Response) => {
                    if (response.ok) {
                        return response.json()
                    }
                    throw "Could not load data"
                })
                .then((data) => {
                    tasks.set(data)
                    error.set(undefined)
                })
                .catch((err) => error.set(err))
        },

        async removeById(taskId: string) {
            tasks.update((tasks: Task[]) => tasks.filter((item) => item.id !== taskId))
        },

        async insert(name: string) {
            const task: Task = { id: crypto.randomUUID(), name: name, complete: false }
            fetch(getJsonRequest(URL, "POST", task))
                .then((response: Response) => {
                    if (response.ok) {
                        return response.json()
                    }
                    throw "Task could not be saved"
                })
                .then((data) => {
                    tasks.update((tasks: Task[]) => [...tasks, data])
                    error.set(undefined)
                })
                .catch((err) => error.set(err))
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
