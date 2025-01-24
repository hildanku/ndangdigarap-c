import { Context } from 'hono'
import { getAll } from './task.service'

export const getAllTask = async (c: Context) => {
    try {
        const tasks = getAll
        
        if (tasks.length === 0) {
            return c.json("task not found")
        }


        return c.json(tasks)
    } catch (error: any) {
        return c.json(error)
    } 
}
