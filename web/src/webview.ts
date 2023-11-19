import { Webview, WebviewPanel, Uri, TextDocument, workspace } from "vscode";

export function getWebviewContent(webview: Webview, extensionUri: Uri) {
  const stylesUri = getUri(webview, extensionUri, [
    "webview",
    "dist",
    "assets",
    "index.css",
  ]);

  const scriptUri = getUri(webview, extensionUri, [
    "webview",
    "dist",
    "assets",
    "index.js",
  ]);

  const codiconsUri = getUri(webview, extensionUri, [
    "webview",
    "dist",
    "codicons",
    "codicon.css",
  ]);

  return /*html*/ `
      <!DOCTYPE html>
      <html lang="en">
        <head>
          <meta charset="UTF-8" />
          <meta name="viewport" content="width=device-width, initial-scale=1.0" />
          <link rel="stylesheet" type="text/css" href="${stylesUri}">
          <link href="${codiconsUri}" rel="stylesheet" />
          <title>Neva Editor</title>
        </head>
        <body>
          <div id="root"></div>
          <script type="module" nonce="${getNonce()}" src="${scriptUri}"></script>
        </body>
      </html>
    `;
}

function getNonce() {
  let text = "";
  const possible =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  for (let i = 0; i < 32; i++) {
    text += possible.charAt(Math.floor(Math.random() * possible.length));
  }
  return text;
}

function getUri(webview: Webview, extensionUri: Uri, pathList: string[]) {
  return webview.asWebviewUri((Uri as any).joinPath(extensionUri, ...pathList));
}

export function sendIndexMsgToWebView(
  panel: WebviewPanel,
  document: TextDocument,
  module: unknown
) {
  panel.webview.postMessage({
    type: "index",
    uri: workspace.workspaceFolders![0].uri,
    document: document,
    module: module,
  });
}

export function sendTabChangeMsgToWebView(
  panel: WebviewPanel,
  document: TextDocument
) {
  panel.webview.postMessage({
    type: "tab_change",
    uri: workspace.workspaceFolders![0].uri,
    document: document,
  });
}