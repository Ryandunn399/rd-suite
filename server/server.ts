
// Import the express in typescript file
import express, { Express, Request, Response } from 'express';
import cors from 'cors' 
import bp from 'body-parser'
import dotenv from 'dotenv'
import { userRouter } from './routes/UserRoutes'

dotenv.config();

/**
 * Initialize the express engine.
 */
const app: Express = express();
 
/**
 * Load port number
 */
const port = process.env.PORT;
 
/**
 * Configure NPM packages
 */
app.use(cors({origin: '*'}))
app.use(bp.json())

/**
 * Will handle the user routing.
 */
app.use('/users', userRouter);
 
// Server setup
app.listen(port, () => {
    console.log(`⚡️[server]: Server is running at http://localhost:${port}`);
});