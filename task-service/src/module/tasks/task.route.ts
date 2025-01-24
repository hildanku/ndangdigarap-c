import { Hono } from 'hono'
import { getAllTask } from './task.controller'

const taskRoutes = new Hono()

taskRoutes.get('/tasks', getAllTask)

export default taskRoutes
