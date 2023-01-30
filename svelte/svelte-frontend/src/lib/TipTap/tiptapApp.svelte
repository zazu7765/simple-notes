<script lang="ts">
    import { onMount, onDestroy } from 'svelte';
      import { writable } from 'svelte/store';
      import { Editor } from '@tiptap/core';
      import StarterKit from '@tiptap/starter-kit';
      
      
      // @todo this throws a 'process is not defined' error in the Svelte REPL.
      // Uncomment the next line to see the REPL issue.
      // import BubbleMenu from '@tiptap/extension-bubble-menu'
      import FixedMenu from './FixedMenu.svelte';
      import {browser} from '$app/environment';
      
      
      export let content = '';
      
      const contentStore = writable(content);
  
      let element;
      let editor:any;
      let bubbleMenu;
  
      onMount(() => {
      editor = new Editor({
        element,
              extensions: [StarterKit],
              editorProps: {
    attributes: {
      class: 'prose prose-sm sm:prose lg:prose-lg xl:prose-2xl m-5 focus:outline-none',
    },
  },
              onTransaction: () => {
          editor = editor;
              },
          });
          editor.on('update', ({ editor }) => {
        console.log('editor html', editor.getHTML());
              contentStore.set(editor.getHTML());
          });
      });
  
      let parser:any;
      onMount(() => {
     
      parser = new DOMParser();
    });

  </script>
  
  <div class="wrapper">
      <FixedMenu {editor} />
      
      <div class="element-wrapper" bind:this={element}/>
  </div>
  
  {#if editor&&browser}

  
  <div >
      {parser.parseFromString(editor.getHTML(), 'text/html').body.firstChild?.textContent }
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