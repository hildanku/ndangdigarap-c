import { Hono } from 'hono'
import { createTask, deleteTask, getAllTask, getTaskById, updateTask } from './task.controller'
import { authMiddleware } from '../middleware/auth'

const taskRoutes = new Hono()

taskRoutes.use("*", authMiddleware)
taskRoutes.get('/tasks', getAllTask)
taskRoutes.get('/task/:id', getTaskById)
taskRoutes.post('/task', createTask)
taskRoutes.put('/task/:id', updateTask)
taskRoutes.delete('/task/:id', deleteTask)

export default taskRoutes
