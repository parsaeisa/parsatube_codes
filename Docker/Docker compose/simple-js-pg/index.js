const { Client } = require('pg');

async function main() {
  const client = new Client({
    user: process.env.POSTGRES_USER,
    host: process.env.POSTGRES_HOST,
    database: process.env.POSTGRES_DB,
    password: process.env.POSTGRES_PASSWORD,
    port: 5432
  });

  try {
    await client.connect();
    console.log('Connected to PostgreSQL');

    const res = await client.query('SELECT NOW() AS now');
    console.log('Now is:', res.rows[0].now);

  } catch (err) {
    console.error('Error:', err);
  } finally {
    await client.end();
  }
}

main();
