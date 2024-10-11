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
            fetch(getJsonRequest(URL + `/${taskId}`, "DELETE"))
                .then((response: Response) => {
                    if (!response.ok) {
                        throw "Task could not be deleted."
                    }
                    tasks.update((tasks: Task[]) => tasks.filter((item) => item.id !== taskId))
                    error.set(undefined)
                })
                .catch((err) => error.set(err))
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

        async update(task: Task) {
            fetch(getJsonRequest(URL + `/${task.id}`, "PUT", task))
                .then((response: Response) => {
                    if (response.ok) {
                        return response.json()
                    }
                    throw "Task could not be saved."
                })
                .then((data: Task) => {
                    tasks.update((tasks: Task[]) => {
                        const taskIndex = tasks.findIndex((task: Task) => task.id == data.id)
                        if (taskIndex >= 0) {
                            tasks[taskIndex] = data
                        }
                        return tasks
                    })
                    error.set(undefined)
                })
                .catch((err) => error.set(err))
        }
    }
}

export default createStore()
