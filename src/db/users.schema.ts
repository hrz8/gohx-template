import {sql} from 'drizzle-orm';
import {pgEnum, pgTable, timestamp, uuid, varchar} from 'drizzle-orm/pg-core';

export const userRole = pgEnum('user_role', ['superadmin', 'member']);

export const users = pgTable('users', {
  id: uuid('id')
    .primaryKey()
    .default(sql`gen_random_uuid()`),
  email: varchar('email').unique(),
  firstName: varchar('first_name').notNull(),
  lastName: varchar('last_name'),
  role: userRole('role').notNull(),
  createdAt: timestamp('created_at', {withTimezone: true}).default(sql`now()`),
  updatedAt: timestamp('updated_at', {withTimezone: true}).default(sql`now()`),
});
