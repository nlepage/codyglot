import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import 'rxjs/add/operator/finally';

export interface ExecuteResult {
  stdout: string;
  stderr: string;
  compilationTime: string;
  runningTime: string;
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
          stdin,
        },
      )
      .finally(() => this.executing = false)
      .subscribe({
        error: (e) => console.error(e), // FIXME do something useful
        next: (result) => this.result = result,
      });
  }
}
