import { integer, pgTable, timestamp, varchar } from 'drizzle-orm/pg-core'
 
export const taskTable = pgTable("tasks", {
    id: integer().primaryKey().generatedAlwaysAsIdentity(),
    user: integer('user').notNull(),
    title: varchar().notNull(),
    description: varchar(),
    status: varchar().default('pending'),
    priority: varchar().default('delete'),
    deadline: varchar(),
    createdAt: timestamp().defaultNow(),
    updatedAt: timestamp().defaultNow()
})
