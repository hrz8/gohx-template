import dotenv from 'dotenv';
dotenv.config();

import {drizzle} from 'drizzle-orm/node-postgres';
import {migrate} from 'drizzle-orm/node-postgres/migrator';
import {Client} from 'pg';

const {DATABASE_URL} = process.env;
const MIGRATION_FOLDER = './scripts/migrations';

if (!DATABASE_URL) {
  throw new Error('cannot resolve database connection!');
}

void (async function (): Promise<void> {
  const client = new Client({
    connectionString: DATABASE_URL,
  });

  await client.connect();
  const db = drizzle(client);

  await migrate(db, {migrationsFolder: MIGRATION_FOLDER});

  await client.end();
})();
