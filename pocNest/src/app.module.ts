import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ScheduleModule } from '@nestjs/schedule';
import { GmailService } from './gmail/gmail.service';
import { GmailController } from './gmail/gmail.controller';

@Module({
  imports: [
    ScheduleModule.forRoot(),
  ],
  controllers: [AppController, GmailController],
  providers: [AppService, GmailService],
})
export class AppModule {}
