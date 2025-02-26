import { Hono } from 'hono'
import taskRoutes from './module/tasks/task.route'
import { logger } from 'hono/logger'


const app = new Hono()
app.use(logger())
app.route('/api', taskRoutes)

export default {
    port: 5010,
    fetch: app.fetch
}
