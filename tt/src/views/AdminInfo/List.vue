<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <div class="table-page-search-wrapper">
        <a-form layout="inline">
          <a-row :gutter="48">
  
          </a-row>
        </a-form>
      </div>

      <div class="table-operator">
		<a-button type="primary" icon="plus"  v-if="$auth('AdminInfo.add')" @click="handleAdd">新建</a-button>
        <a-dropdown v-action:edit v-if="selectedRowKeys.length > 0">
          <a-menu slot="overlay">
            <a-menu-item key="1"  v-if="$auth('AdminInfo.delBatch')"><a-icon type="delete" />删除</a-menu-item>
          </a-menu>
          <a-button style="margin-left: 8px">
            批量操作 <a-icon type="down" />
          </a-button>
        </a-dropdown>
      </div>

      <s-table
              ref="table"
              size="default"
              rowKey="key"
              :columns="columns"
              :data="loadData"
              :alert="true"
              :rowSelection="rowSelection"
              showPagination="auto"
      >
        <span slot="serial" slot-scope="text, record, index">
          {{ index + 1 }}
        </span>
        <span slot="action" slot-scope="text, record">
          <template>
		<a  v-if="$auth('AdminInfo.edit')" @click="handleEdit(record)">编辑</a>
		<a-divider  v-if="$auth('AdminInfo.edit')" type="vertical" />
		<a v-if="$auth('AdminInfo.del')" @click="handleDelete(record.id)">删除</a>
          </template>
        </span>
      </s-table>

    </a-card>
  </page-header-wrapper>
</template>

<script>
import moment from 'moment'
import { STable } from '@/components'
import { list } from '@/api/AdminInfo'

const columns = [
	{
		title:'',
		dataIndex:'nick_name',
		key:'nick_name',
	}
]

export default {
  name: 'TableList',
  components: {
    STable,
  },
  data () {
    this.columns = columns
    return {
              // create model
              confirmLoading: false,
              // 高级搜索 展开/关闭
              advanced: false,
              // 查询参数
              queryParam: {},
              // 加载数据方法 必须为 Promise 对象
              loadData: parameter => {
                const requestParameters = Object.assign({}, parameter, this.queryParam)
                console.log('loadData request parameters:', requestParameters)
                return getList(requestParameters)
                        .then(res => {
                          console.log(res)
                          return res.result
                        })
              },

            selectedRowKeys: [],
            selectedRows: []
    }
  },

  computed: {
    rowSelection () {
      return {
        selectedRowKeys: this.selectedRowKeys,
        onChange: this.onSelectChange
      }
    }
  },
  methods: {
    handleDelete (id) {
      const _this = this
      this.$confirm({
        title: '你确定删除这个选项吗?',
        content: '删除后将不能恢复，我们将记录您的炒作行为！',
        okText: 'Yes',
        okType: 'danger',
        cancelText: 'No',
        onOk () {
          del(id).then(res => {
            _this.getList()
          })
        }
      })
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    resetSearchForm () {
      this.queryParam = {
        date: moment(new Date())
      }
    }
  }
}
</script>
