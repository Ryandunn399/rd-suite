// Import the express in typescript file
import  { Router, Express, Request, Response } from 'express';

export const userRouter = Router();

/**
 * The type for our registration response object.
 */
interface RegisterResponse {
    user: string,
    id: string
}

/**
 * Unused test route to ensure router is working via browser.
 */
userRouter.get('/register', (req: Request, res: Response) => {
    res.send("Hello!")   
})

/**
 * Our register route that currently has a hard coded value to return to the user.
 */
userRouter.post('/register', async (req: Request, res: Response) => {
    let responseObj: RegisterResponse;

    responseObj = {
        user: 'Ryan',
        id: '1'
    }

    res.send(JSON.stringify(responseObj))
})