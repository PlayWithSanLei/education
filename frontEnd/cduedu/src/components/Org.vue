<template>
  <div class="custom-tree-container">
    <el-card class="card">
      <el-tree
          :data="data"
          node-key="id"
          default-expand-all
          :expand-on-click-node="false"
      >
        <template #default="{ node, data }">
        <span class="custom-tree-node">
          <span>{{ node.label,data}}</span>
          <span>
            <a @click="append(data)"> Append </a>
            <a @click="remove(node, data)"> Delete </a>
          </span>
        </span>
        </template>
      </el-tree>
    </el-card>

  </div>
</template>

<script>
let id = 1000

export default {
  name: 'Org',
  data() {
    const data = [
      {
        id: 1,
        label: '一级 1',
        children: [
          {
            id: 4,
            label: '二级 1-1',
            children: [
              {
                id: 9,
                label: '三级 1-1-1',
              },
              {
                id: 10,
                label: '三级 1-1-2',
              },
            ],
          },
        ],
      },
      {
        id: 2,
        label: '一级 2',
        children: [
          {
            id: 5,
            label: '二级 2-1',
          },
          {
            id: 6,
            label: '二级 2-2',
          },
        ],
      },
      {
        id: 3,
        label: '一级 3',
        children: [
          {
            id: 7,
            label: '二级 3-1',
          },
          {
            id: 8,
            label: '二级 3-2',
          },
        ],
      },
    ]
    return {
      data: JSON.parse(JSON.stringify(data)),
    }
  },

  methods: {
    append(data) {
      const newChild = {id: id++, label: 'testtest', children: []}
      if (!data.children) {
        data.children = []
      }
      data.children.push(newChild)
      this.data = [...this.data]
    },

    remove(node, data) {
      const parent = node.parent
      const children = parent.data.children || parent.data
      const index = children.findIndex((d) => d.id === data.id)
      children.splice(index, 1)
      this.data = [...this.data]
    },

    renderContent(h, {node, data, store}) {
      return h(
          'span',
          {
            class: 'custom-tree-node',
          },
          h('span', null, node.label),
          h(
              'span',
              null,
              h(
                  'a',
                  {
                    onClick: () => this.append(data),
                  },
                  'Append '
              ),
              h(
                  'a',
                  {
                    onClick: () => this.remove(node, data),
                  },
                  'Delete'
              )
          )
      )
    },
  },
}
</script>
<style>
.custom-tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 14px;
  padding-right: 8px;
}

.card {
  margin: 1em 0 0 0;
  box-shadow: -1px 1px 20px grey;
  height: 95%;
}
</style>
