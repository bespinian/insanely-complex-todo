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
            try {
                await fetch(getJsonRequest(URL + `/${taskId}`, "DELETE"))
                tasks.update((tasks: Task[]) => tasks.filter((item) => item.id !== taskId))
                error.set(undefined)
            } catch {
                error.set("Could not delete task.")
            }
        },

        async insert(name: string) {
            const task: Task = { id: crypto.randomUUID(), name: name, complete: false }
            try {
                const jsonResponse = await fetch(getJsonRequest(URL, "POST", task))
                const data = await jsonResponse.json()
                tasks.update((tasks: Task[]) => [...tasks, data])
                error.set(undefined)
            } catch {
                error.set("Could not add task.")
            }
        },

        async update(task: Task) {
            try {
                const jsonResponse = await fetch(getJsonRequest(URL + `/${task.id}`, "PUT", task))
                const data: Task = await jsonResponse.json()

                tasks.update((tasks: Task[]) => {
                    const taskIndex = tasks.findIndex((task: Task) => task.id == data.id)
                    if (taskIndex >= 0) {
                        tasks[taskIndex] = data
                    }
                    return tasks
                })
                error.set(undefined)
            } catch {
                error.set("Could not update task.")
            }
        }
    }
}

export default createStore()
