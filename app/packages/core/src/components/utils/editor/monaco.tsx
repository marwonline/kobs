import { loader } from '@monaco-editor/react';
import { Monaco } from '@monaco-editor/react';
import * as monaco from 'monaco-editor';
import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker';
import jsonWorker from 'monaco-editor/esm/vs/language/json/json.worker?worker';
import { promLanguageDefinition } from 'monaco-promql';
import { setDiagnosticsOptions } from 'monaco-yaml';

import CompletionItem = monaco.languages.CompletionItem;
import ProviderResult = monaco.languages.ProviderResult;
import CompletionList = monaco.languages.CompletionList;

import yamlWorker from './yaml.worker.js?worker';

/**
 * `setupMonaco` is used to setup monaco, which is used in various places of the app, e.g. to show and edit yaml files
 * of Kubernetes resources or to provide autocompletion in plugins for the query language of a service (e.g. PromQL for
 * the Prometheus Plugin).
 */
export const setupMonaco = () => {
  // eslint-disable-next-line no-restricted-globals
  self.MonacoEnvironment = {
    getWorker(_, label) {
      switch (label) {
        case 'editorWorkerService':
          return new editorWorker();
        case 'json':
          return new jsonWorker();
        case 'yaml':
          return new yamlWorker();
        default:
          throw new Error(`Unknown label ${label}`);
      }
    },
  };

  loader.config({ monaco });

  loader.init().then(() => {
    // Initialized
  });

  /**
   * `setDiagnosticsOptions` is used to setup the `monaco-yaml` plugin. This would also allow us to register custom
   * schemas later.
   */
  setDiagnosticsOptions({
    completion: true,
    enableSchemaRequest: true,
    format: true,
    hover: true,
    validate: true,
  });
};

/**
 * `PROMQL_SETUP_STARTED` is a constants which allows us to check if the `setupPromQL` functions was already run, so
 * that the setup function is only run once.
 */
let PROMQL_SETUP_STARTED = false;

/**
 * `setupPromQL` runs the setup for PromQL support. It adds syntax highlighting and autocompletion via the
 * `monaco-promql` plugin. It is also possible to load custom completion items via the passed in `loadCompletionItems`
 * function.
 */
export const setupPromQL = (monaco: Monaco, loadCompletionItems?: () => Promise<string[]>) => {
  if (PROMQL_SETUP_STARTED === false) {
    PROMQL_SETUP_STARTED = true;

    const { aliases, extensions, mimetypes, loader } = promLanguageDefinition;
    monaco.languages.register({ aliases, extensions, id: promLanguageDefinition.id, mimetypes });

    loader().then((mod) => {
      monaco.languages.setMonarchTokensProvider(promLanguageDefinition.id, mod.language);
      monaco.languages.setLanguageConfiguration(promLanguageDefinition.id, mod.languageConfiguration);
      monaco.languages.registerCompletionItemProvider(promLanguageDefinition.id, mod.completionItemProvider);
      monaco.languages.registerCompletionItemProvider(promLanguageDefinition.id, {
        provideCompletionItems: async (model, position, context, token) => {
          if (loadCompletionItems) {
            const labels = await loadCompletionItems();

            const suggestions = labels.map((value) => {
              return {
                insertText: value,
                insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
                kind: monaco.languages.CompletionItemKind.Keyword,
                label: value,
              } as CompletionItem;
            });

            return { suggestions } as ProviderResult<CompletionList>;
          }
        },
      });
    });
  }
};
