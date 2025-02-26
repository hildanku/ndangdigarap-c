import { eq } from "drizzle-orm";
import { db } from "../../db/config";
import { taskTable, type Task } from "../../db/schema"

export class TaskRepository {

    constructor(
        private dbInstance = db
    ) { }

    async getAll() {
        return await this.dbInstance
            .select()
            .from(taskTable)
    }

    async getById(id: number) {
        return await this.dbInstance
            .select()
            .from(taskTable)
            .where(eq(taskTable.id, id))
    }

    async getByUserID(user: number) {
        return await this.dbInstance
            .select()
            .from(taskTable)
            .where(eq(taskTable.user, user))
    }

    async create(task: Task) {
        return await this.dbInstance
            .insert(taskTable)
            .values(task)
            .returning()
    }

    async update(id: number, task: Partial<Task>) {
        return this.dbInstance
            .update(taskTable)
            .set(task)
            .where(
                eq(
                    taskTable.id, id
                )
            )
            .returning()
    }

    async delete(id: number) {
        return this.dbInstance
            .delete(taskTable)
            .where(
                eq(
                    taskTable.id, id
                )
            )
            .returning()
    }
}
