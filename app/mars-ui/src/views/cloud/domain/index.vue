<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-row :gutter="24">
          <el-col :span="24" :xs="24">
            <el-col :span="24" :xs="24">
              <el-row :gutter="10" class="mb8">
                <el-col :span="1.5">
                  <el-button
                    v-permisaction="['cloud:domain:add']"
                    type="primary"
                    icon="el-icon-plus"
                    size="mini"
                    @click="handlerCreate"
                  >创建云主机</el-button>
                </el-col>
                <el-col :span="1.5">
                  <el-button
                    v-permisaction="['cloud:domain:edit']"
                    type="success"
                    icon="el-icon-edit"
                    size="mini"
                    :disabled="single"
                    @click="handleUpdate"
                  >修改</el-button>
                </el-col>
                <el-col :span="1.5">
                  <el-button
                    v-permisaction="['cloud:domain:remove']"
                    type="danger"
                    icon="el-icon-delete"
                    size="mini"
                    :disabled="multiple"
                    @click="handleDelete"
                  >删除</el-button>
                </el-col>
              </el-row>

              <el-table
                v-loading="loading"
                :data="domainList"
                border
              >
                <el-table-column type="selection" align="center" />
                <el-table-column label="编号" prop="domainId" sortable="custom" />
                <el-table-column label="主机名" prop="name" sortable="custom" :show-overflow-tooltip="true" />
                <el-table-column label="核心数(VCPU)" prop="vCpuNum" sortable="custom" :show-overflow-tooltip="true" />
                <el-table-column label="内存(GB)" prop="memory" sortable="custom" :show-overflow-tooltip="true" />
                <el-table-column
                  label="操作"
                  fix="right"
                  class-name="small-padding fixed-width"
                >
                  <template slot-scope="scope">
                    <el-button
                      v-permisaction="['cloud:domain:edit']"
                      size="mini"
                      type="text"
                      icon="el-icon-edit"
                      @click="handleUpdate(scope.row)"
                    >修改</el-button>
                    <el-button
                      v-if="scope.row.userId !== 1"
                      v-permisaction="['acloud:domain:remove']"
                      size="mini"
                      type="text"
                      icon="el-icon-delete"
                      @click="handleDelete(scope.row)"
                    >删除</el-button>
                    <el-button
                      v-permisaction="['cloud:domain:reset']"
                      size="mini"
                      type="text"
                      icon="el-icon-key"
                      @click="handleReset(scope.row)"
                    >重置</el-button>
                  </template>
                </el-table-column>
              </el-table>
              <pagination
                v-show="total>0"
                :total="total"
                :page.sync="queryParams.pageIndex"
                :limit.sync="queryParams.pageSize"
                @pagination="getList"
              />
            </el-col>
          </el-col>
        </el-row>
      </el-card>
      <el-dialog :title="title" :visible.sync="open" width="600px">
        <el-form ref="form" :model="form" :rules="rules" label-width="80px">
          <el-row>
            <el-col :span="24">
              <el-form-item label="名称" prop="name">
                <el-input v-model="form.name" placeholder="请输入云主机名称" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="核心数" prop="vCpuNum">
                <el-select v-model="form.vCpuNum" placeholder="请选择" @change="$forceUpdate()">
                  <el-option
                    v-for="dict in cpuOptions"
                    :key="dict.value"
                    :label="dict.label"
                    :value="dict.value"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="内存" prop="memory">
                <el-select v-model="form.memory" placeholder="请选择" @change="$forceUpdate()">
                  <el-option
                    v-for="dict in memoryOptions"
                    :key="dict.value"
                    :label="dict.label"
                    :value="dict.value"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="描述">
                <el-input v-model="form.description" type="textarea" placeholder="请输入内容" />
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button type="primary" @click="submitForm">确 定</el-button>
          <el-button @click="cancel">取 消</el-button>
        </div>
      </el-dialog>

    </template>
  </BasicLayout>
</template>

<script>
import { listDomain } from '@/api/cloud/domain'
export default {
  name: 'DomainManage',
  components: {},
  data() {
    return {
      // 遮罩层
      loading: false,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 云主机表格数据
      domainList: null,
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      // 日期范围
      dateRange: [],
      // 状态数据字典
      statusOptions: [],
      // cpu数据字典
      cpuOptions: [],
      // 内存数据字典
      memoryOptions: [],
      // 镜像选项
      imageOptions: [],
      // 表单参数
      form: {},
      defaultProps: {
        children: 'children',
        label: 'label'
      },
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        status: undefined
      },
      rules: {
        name: [{ required: true, message: '云主机名称不能为空', trigger: 'blur' }],
        memory: [{ required: true, message: '内存不能为空', trigger: 'blur' }],
        vCpuNum: [{ required: true, message: '用户密码不能为空', trigger: 'blur' }]
      }
    }
  },
  created() {
    this.getList()
    this.getDicts('cloud_cpu_num').then(response => {
      this.cpuOptions = response.data
    })
    this.getDicts('cloud_domain_memory').then(response => {
      this.memoryOptions = response.data
    })
  },
  methods: {
    getList() {
      // this.loading = true
      listDomain(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.domainList = response.data.list
        this.total = response.data.count
        this.loading = false
      }
      )
    },
    handlerCreate() {
      this.open = true
      this.title = '创建云主机'
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {
        domainId: undefined,
        name: undefined,
        memory: undefined,
        vCpuNum: undefined,
        description: undefined
      }
      this.resetForm('form')
    },
    handleDelete() {},
    handleUpdate() {},
    handleReset() {},
    submitForm() {}
  }
}

</script>

<style scoped>

</style>
