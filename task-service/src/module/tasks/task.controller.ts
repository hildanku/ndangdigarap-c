import { Context } from 'hono'
import { TaskService } from './task.service'
import { TaskRepository } from './task.repository'
import { Task } from '../../db/schema'
import { appResponse } from '../../helpers/response'

const taskRepo = new TaskRepository()
const taskSvc = new TaskService(taskRepo)

export const getAllTask = async (c: Context) => {
    try {
        const tasks = await taskSvc.getAll()
        if (tasks.length === 0) {
            return appResponse(c, 404, 'task not found', null)
        }
        return appResponse(c, 200, 'tasks found', tasks)
    } catch (error: any) {
        return appResponse(c, 500, 'error', null)
    }
}

export const getTaskById = async (c: Context) => {
    const id = Number(c.req.param('id'))
    try {
        const task = await taskSvc.getById(id)

        if (task.length === 0) {
            return appResponse(c, 404, 'task not found', null)
        }
        return appResponse(c, 200, 'task found', task[0])
    } catch (error: any) {
        return appResponse(c, 500, 'something when wrong', null)
    }
}

export const createTask = async (c: Context) => {
    const task: Task = await c.req.json()

    const u = c.get('user')
    task.user = u.user_id
    console.log(u)

    try {
        const newTask = await taskSvc.create(task)
        return appResponse(c, 201, 'task created', newTask)
    } catch (error: any) {
        return appResponse(c, 500, 'something when wrong', null)
    }
}

export const updateTask = async (c: Context) => {
    const id = Number(c.req.param('id'))
    const task: Partial<Task> = await c.req.json()
    try {
        const updatedTask = await taskSvc.update(id, task)
        if (updatedTask.length === 0) {
            return appResponse(c, 404, 'task not found', null)
        }
        return appResponse(c, 200, 'task updated', updatedTask[0])
    } catch (error: any) {
        return appResponse(c, 500, 'something when wrong', null)
    }
}

export const deleteTask = async (c: Context) => {
    const id = Number(c.req.param('id'))
    try {
        const deletedTask = await taskSvc.delete(id)
        if (deletedTask.length === 0) {
            return appResponse(c, 404, 'task not found', null)
        }
        return appResponse(c, 200, 'task deleted', null)
    } catch (error: any) {
        return appResponse(c, 500, 'something when wrong', null)
    }
}
