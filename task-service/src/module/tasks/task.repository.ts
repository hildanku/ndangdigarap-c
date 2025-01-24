import { eq } from "drizzle-orm";
import { db } from "../../db/config";
import { taskTable } from "../../db/schema";

export class TaskRepository {
    async getAll() {
        const tasks = await db
            .select()
            .from(taskTable)

        return tasks
    }

    async getByUserID(user: number) {
        const task = await db
            .select()
            .from(taskTable)
            .where(eq(taskTable.user, user))
        
        return task
    }

}
