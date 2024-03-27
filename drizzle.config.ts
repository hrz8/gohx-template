import dotenv from 'dotenv';
dotenv.config();

import type {Config} from 'drizzle-kit';

const {DATABASE_URL} = process.env;
const MIGRATION_FOLDER = './scripts/migrations';

if (!DATABASE_URL) {
  throw new Error('cannot resolve database connection!');
}

export default {
  schema: './src/db/**/*.schema.ts',
  out: MIGRATION_FOLDER,
  driver: 'pg',
  dbCredentials: {connectionString: DATABASE_URL},
} satisfies Config;
