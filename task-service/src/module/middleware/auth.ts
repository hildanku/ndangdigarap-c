import { Context, Next } from "hono"
import { appResponse } from "../../helpers/response"
import { verify } from "hono/jwt"

export const authMiddleware = async (c: Context, next: Next) => {
    const authHeader = c.req.header('Authorization')
    if (!authHeader) {
        return appResponse(c, 401, 'unauthorized', null)
    }

    const token = authHeader.split(' ')[1]
    try {
        const payload = await verify(token, 'hatsunemikuluplup')
        c.set('user', payload)
        await next()
    } catch (error) {
        return appResponse(c, 401, 'unauthorized', null)
    }
}
