import { Hono } from 'hono'
import taskRoutes from './module/tasks/task.route'

const app = new Hono()

app.route('/api', taskRoutes)

export default {
    port: 5010,
    fetch: app.fetch
}
