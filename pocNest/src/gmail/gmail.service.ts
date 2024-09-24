import { Injectable } from '@nestjs/common';
import { Auth, google } from 'googleapis';
import 'dotenv/config';

@Injectable()
export class GmailService {
    private oAuth2Client: Auth.OAuth2Client;

    constructor() {
        this.oAuth2Client = new google.auth.OAuth2(
            `${process.env.GOOGLE_CLIENT_ID}`,
            `${process.env.GOOGLE_CLIENT_SECRET}`,
            `${process.env.GOOGLE_REDIRECT_URI}`,
        );
    }

    generateAuthUrl() {
        return this.oAuth2Client.generateAuthUrl({
            access_type: 'offline',
            scope: [
                'https://www.googleapis.com/auth/gmail.modify',
                'https://www.googleapis.com/auth/gmail.send',
                'email',
                'profile'
            ],
            prompt: 'consent'
        });
    }

    async getTokens(code: string) {
        const { tokens } = await this.oAuth2Client.getToken(code);
        this.oAuth2Client.setCredentials(tokens);
        return tokens;
    }

    getOAuth2Client() {
        return this.oAuth2Client;
    }
}
