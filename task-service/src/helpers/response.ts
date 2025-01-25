import { Context } from "hono";

export const appResponse = (c: Context, status: number, message: string, results: any) => {
    const responseData: responseData = {
        message,
        results,
    }
    return c.json(responseData, status)
}

/* usage
 * app.get('/example', (c) => {
    const data = { id: 1, name: 'John Doe' };
    return appResponse(c, 200, 'Success', data);
});
 *
 *
 */
