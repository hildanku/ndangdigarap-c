import { Context } from "hono";
import { ContentfulStatusCode } from "hono/utils/http-status";

export const appResponse = (c: Context, status: ContentfulStatusCode, message: string, results: any) => {
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
