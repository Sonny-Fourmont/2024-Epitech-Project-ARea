import { Controller, Get, Query, Redirect, Res, Logger } from '@nestjs/common';
import { Cron, CronExpression } from '@nestjs/schedule';
import { GmailService } from './gmail.service';
import { google } from 'googleapis';
import 'dotenv/config';

@Controller('gmail')
export class GmailController {
    constructor(private readonly gmailServices: GmailService) {};
    private readonly logger: any = new Logger(GmailController.name)
    private accessToken: string;

    @Get()
    @Redirect()
    googlelogin() {
      const url: string = this.gmailServices.generateAuthUrl();
      return {url};
    }

    @Get('callback')
    async googleCallback(@Query('code') code: string, @Res() res: any) {
      const tokens = await this.gmailServices.getTokens(code);
      if (tokens.access_token) {
        this.accessToken = tokens.access_token
        this.logger.log(`The access token has been given`);
      }
      res.json(tokens);
    }

    @Cron(CronExpression.EVERY_MINUTE)
    async CronJob() {
      const authClient: any = this.gmailServices.getOAuth2Client();
      const gmail: any = google.gmail({version: 'v1', auth: authClient});
      var email: string;
      var base64Email: string;

      if (this.accessToken) {
        const mailDocument = [
          `From: ${gmail.user}`,
          `To: ${process.env.DEST_EMAIL}`,
          'Content-type: text/html;charset=iso-8859-1',
          'MIME-Version: 1.0',
          'Subject: AREA Test',
          '',
          'This is a test email'
        ];

        email = mailDocument.join('\r\n').trim();
        base64Email = Buffer.from(email).toString('base64');

        await gmail.users.messages.send({
          userId: 'me',
          requestBody: {
            raw: base64Email
          }
        });

        this.logger.log('mail sent !');
      } else {
        this.logger.error('Token not given');
      }
    }
}
