import { HttpClient, HttpFetchBackend, setupConfig } from "@ngify/http";
import { catchError, finalize, tap, throwError } from "rxjs";
import { getToken } from "@/utils/auth";

setupConfig({
  backend: new HttpFetchBackend(),
});

const request = new HttpClient([
  {
    intercept(req, next) {
      // Clone the request to modify request parameters
      const token = getToken();
      req = req.clone({
        url: `${import.meta.env.VITE_API_URL}${req.url}`,
      });
      if (req.method === "POST") {
        req = req.clone({
          headers: req.headers.set("Content-Type", `application/json`),
        });
      }
      if (token) {
        req = req.clone({
          headers: req.headers.set("Authorization", `bearer ${token}`),
        });
      }

      return next.handle(req);
    },
  },
  {
    intercept(req, next) {
      const started = Date.now();
      let ok;
      return next.handle(req).pipe(
        tap(
          // Succeeds when there is a response; ignore other events
          (res) => {
            ok = "success";
            const body = res.body;
            if (!body.success) {
              throw new Error(JSON.stringify(body));
            }
          },
          // Operation failed; error is an HttpErrorResponse
          () => (ok = "failed")
        ),
        catchError((err) => {
          return throwError(err);
        }),
        finalize(() => {
          const elapsed = Date.now() - started;
          const msg = `${req.method} "${req.urlWithParams}" ${ok} in ${elapsed} ms.`;
          console.log(msg);
        })
      );
    },
  },
]);

export { request };
