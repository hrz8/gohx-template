import {relations, sql} from 'drizzle-orm';
import {
  jsonb,
  pgTable,
  text,
  timestamp,
  uuid,
  varchar,
} from 'drizzle-orm/pg-core';

import {users} from './users.schema';

export type AppSettings = {
  incomingURL: string;
};

export const apps = pgTable('apps', {
  id: uuid('id')
    .primaryKey()
    .default(sql`gen_random_uuid()`),
  ownerId: uuid('owner_id')
    .notNull()
    .references(() => users.id, {onDelete: 'cascade'}),
  name: varchar('name', {length: 25}).notNull(),
  subtitle: varchar('subtitle', {length: 75}),
  description: text('description'),
  settings: jsonb('settings')
    .$type<AppSettings>()
    .default(sql`'{}'::jsonb`),
  createdAt: timestamp('created_at', {withTimezone: true}).default(sql`now()`),
  updatedAt: timestamp('updated_at', {withTimezone: true}).default(sql`now()`),
});

export const appRelations = relations(apps, ({one}) => ({
  owner: one(users, {
    fields: [apps.ownerId],
    references: [users.id],
  }),
}));
