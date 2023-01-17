const chainer = (editor, fn) => editor.chain().focus()[fn]().run();

const toggleBold = editor => chainer(editor, 'toggleBold');
const toggleItalic = editor => editor.chain().focus().toggleItalic().run();
const toggleStrike = editor => editor.chain().focus().toggleStrike().run();

export { toggleBold, toggleItalic, toggleStrike };