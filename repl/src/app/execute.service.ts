import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Subject } from 'rxjs';

type ExecuteResult = {
  stdout: string,
  stderr: string,
}

@Injectable()
export class ExecuteService {

  private _result = new Subject<ExecuteResult>();

  constructor(private http: HttpClient) {}

  execute = (language: string, source: string, stdin: string): void => {
    this.http.post<ExecuteResult>('/api/execute', { language, source, stdin })
      .subscribe({
        next: result => this._result.next(result),
        error: e => this._result.error(e),
      });
  }

  get result() { return this._result.asObservable() }

}
