<script lang="ts">
    import { onMount, onDestroy } from 'svelte';
      import { writable } from 'svelte/store';
      import { Editor } from '@tiptap/core';
      import StarterKit from '@tiptap/starter-kit';
      
      // @todo this throws a 'process is not defined' error in the Svelte REPL.
      // Uncomment the next line to see the REPL issue.
      // import BubbleMenu from '@tiptap/extension-bubble-menu'
      import FixedMenu from './FixedMenu.svelte';
      
      export let content = '';
      
      const contentStore = writable(content);
  
      let element;
      let editor:any;
      let bubbleMenu;
  
      onMount(() => {
      editor = new Editor({
        element,
              extensions: [StarterKit],
            
              onTransaction: () => {
          editor = editor;
              },
          });
          editor.on('update', ({ editor }) => {
        console.log('editor html', editor.getHTML());
              contentStore.set(editor.getHTML());
          });
      });
  
      onDestroy(() => {
      editor.destroy();
      });
      const parser = new DOMParser();

  </script>
  
  <div class="wrapper">
      <FixedMenu {editor} />
      
      <div class="element-wrapper" bind:this={element}/>
  </div>
  
  {#if editor}

  
  <div >
      {parser.parseFromString(editor.getHTML(), 'text/html').body.firstChild?.textContent}
  </div>

  

  {/if}
  
  <style>
      .wrapper {
      border: 1px solid #ccc;
          max-height: 200px;
          display: inline-flex;
          flex-direction: column;
      }
      
      .wrapper:focus-within {
      border: 1px solid red;
      }
      
      .element-wrapper {
      padding: 1rem;
          flex: 1 1 0%;
          resize: both;
          overflow: auto;
      }
      
      .element-wrapper :global(p:first-of-type) {
      margin-top: 0;
      }
      
      .element-wrapper :global(p:last-of-type) {
      margin-bottom: 0;
      }
      
      .element-wrapper > :global(.ProseMirror) {
      outline: 0;
      }
      
      .json-output,
      .html-output {
      max-height: 200px;
          overflow: hidden;
          overflow-y: auto;
          border: 1px solid #ccc;
      }
  </style>