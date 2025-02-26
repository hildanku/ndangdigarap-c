import { Task } from "../../db/schema";
import { appResponse } from "../../helpers/response"
import { TaskRepository } from "./task.repository"

export class TaskService {

    constructor(private taskRepo = new TaskRepository()) { }

    async getAll() {
        return await this.taskRepo.getAll()
    }

    async getById(id: number) {
        return await this.taskRepo.getById(id)
    }

    async create(task: Task) {
        return await this.taskRepo.create(task);
    }

    async update(id: number, task: Partial<Task>) {
        return await this.taskRepo.update(id, task);
    }

    async delete(id: number) {
        return await this.taskRepo.delete(id);
    }

}
