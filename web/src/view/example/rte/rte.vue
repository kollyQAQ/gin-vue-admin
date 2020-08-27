<template>
  <div>
    <div class="edit_container">
      <quill-editor
              :options="editorOption"
              @blur="onEditorBlur($event)"
              @change="onEditorChange($event)"
              @focus="onEditorFocus($event)"
              ref="myQuillEditor"
              v-model="content"
      ></quill-editor>
      <el-button plain @click="saveHtml">保存</el-button>
    </div>
    <div class="edit">
      <mavon-editor v-model="model.content"></mavon-editor>
      <el-button type="primary" size="small" @click="submit">发表</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      // content: `<p>hello world</p>`,
      content: `<h2>hello world</h2><p><strong>woshi</strong>yigedashazz</p><blockquote>hhhhhhhh</blockquote><ul><li>111</li><li>222</li><li>333</li></ul>`,
      editorOption: {},
      model:{
        content:''
      }
    }
  },
  computed: {
    editor() {
      return this.$refs.myQuillEditor.quill
    }
  },
  methods: {
    // onEditorReady(editor) { // 准备编辑器

    // },
    onEditorBlur() {}, // 失去焦点事件
    onEditorFocus() {}, // 获得焦点事件
    onEditorChange() {}, // 内容改变事件
    saveHtml() {
      console.log(2222)
      console.log(this.content)
    }, // 保存方法
    submit(){}
  }
}
</script>
