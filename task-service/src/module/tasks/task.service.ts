import { TaskRepository } from "./task.repository"
//import * as taskRepository from './task.repository'

const TaskRepo = new TaskRepository()

export const getAll = async () => {
    return await TaskRepo.getAll()
}
