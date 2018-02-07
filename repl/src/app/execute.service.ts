import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Subject } from 'rxjs/Subject';

export interface ExecuteResult {
  stdout: string;
  stderr: string;
  compilationTime: string;
  runningTime: string;
}

@Injectable()
export class ExecuteService {

  private result$ = new Subject<ExecuteResult>();

  constructor(private http: HttpClient) {}

  execute(language: string, source: string, stdin: string): void {
    this.http.post<ExecuteResult>('/api/execute', { language, source, stdin })
      .subscribe({
        error: (e) => this.result$.error(e),
        next: (result) => this.result$.next(result),
      });
  }

  get result() { return this.result$.asObservable(); }

}
