import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import 'rxjs/add/operator/finally';

interface CommandResult {
  status: number;
  stdout: string;
  stderr: string;
  duration: number;
}

export interface ExecuteResult {
  compilation: CommandResult;
  executions: [CommandResult];
}

@Injectable()
export class ExecuteService {

  result: ExecuteResult;
  executing = false;

  constructor(private http: HttpClient) {}

  execute(language: string, source: string, stdin: string): void {
    this.result = undefined;
    this.executing = true;
    this.http.post<ExecuteResult>(
        '/api/execute',
        {
          language,
          sources: [
            {
              path: {
                golang: 'main.go',
                javascript: 'index.js',
                typescript: 'index.ts',
              }[language],
              content: source,
            },
          ],
          executions: [
            { stdin },
          ],
        },
      )
      .finally(() => this.executing = false)
      .subscribe({
        error: (e) => console.error(e), // FIXME do something useful
        next: (result) => this.result = result,
      });
  }
}
