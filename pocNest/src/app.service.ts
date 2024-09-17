import { Injectable, Logger } from '@nestjs/common';
import { Cron, CronExpression } from '@nestjs/schedule';

@Injectable()
export class AppService {
  private readonly logger = new Logger(AppService.name)

  @Cron(CronExpression.EVERY_MINUTE)
  handleCron() {
    this.logger.debug('called every minutes')
  }

  getHello(): string {
    return 'Hello World!';
  }
}
