<template>
  <div class="connections">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>连接管理</span>
          <el-button type="primary" @click="addConnection">
            <el-icon><Plus /></el-icon>
            添加连接
          </el-button>
        </div>
      </template>
      
      <el-table :data="connections" style="width: 100%">
        <el-table-column prop="name" label="名称" width="200" />
        <el-table-column prop="serverAddr" label="服务器地址" width="200" />
        <el-table-column prop="status" label="状态" width="120">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="180" />
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button 
              size="small" 
              @click="editConnection(scope.row)"
            >
              编辑
            </el-button>
            <el-button 
              size="small" 
              type="danger" 
              @click="deleteConnection(scope.row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    
    <!-- 添加/编辑连接对话框 -->
    <el-dialog 
      v-model="dialogVisible" 
      :title="dialogTitle"
      width="500px"
    >
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="连接名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入连接名称" />
        </el-form-item>
        <el-form-item label="服务器地址" prop="serverAddr">
          <el-input v-model="form.serverAddr" placeholder="例如: 127.0.0.1:10001" />
        </el-form-item>
        <el-form-item label="客户端ID" prop="clientId">
          <el-input v-model="form.clientId" placeholder="自动生成" />
        </el-form-item>
      </el-form>
      
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveConnection">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'

const connections = ref([])
const dialogVisible = ref(false)
const dialogTitle = ref('添加连接')
const formRef = ref()

const form = reactive({
  id: null,
  name: '',
  serverAddr: '',
  clientId: ''
})

const rules = {
  name: [
    { required: true, message: '请输入连接名称', trigger: 'blur' }
  ],
  serverAddr: [
    { required: true, message: '请输入服务器地址', trigger: 'blur' }
  ]
}

const getStatusType = (status) => {
  switch (status) {
    case 'connected':
      return 'success'
    case 'connecting':
      return 'warning'
    case 'error':
      return 'danger'
    default:
      return 'info'
  }
}

const getStatusText = (status) => {
  switch (status) {
    case 'connected':
      return '已连接'
    case 'connecting':
      return '连接中'
    case 'error':
      return '连接错误'
    default:
      return '未连接'
  }
}

const addConnection = () => {
  dialogTitle.value = '添加连接'
  resetForm()
  dialogVisible.value = true
}

const editConnection = (row) => {
  dialogTitle.value = '编辑连接'
  form.id = row.id
  form.name = row.name
  form.serverAddr = row.serverAddr
  form.clientId = row.clientId
  dialogVisible.value = true
}

const deleteConnection = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除这个连接吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    const index = connections.value.findIndex(c => c.id === row.id)
    if (index > -1) {
      connections.value.splice(index, 1)
      ElMessage.success('删除成功')
    }
  } catch {
    // 用户取消删除
  }
}

const saveConnection = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    
    if (form.id) {
      // 编辑
      const index = connections.value.findIndex(c => c.id === form.id)
      if (index > -1) {
        connections.value[index] = {
          ...form,
          status: 'disconnected',
          createTime: connections.value[index].createTime
        }
      }
      ElMessage.success('更新成功')
    } else {
      // 添加
      const newConnection = {
        id: Date.now(),
        ...form,
        status: 'disconnected',
        createTime: new Date().toLocaleString()
      }
      connections.value.push(newConnection)
      ElMessage.success('添加成功')
    }
    
    dialogVisible.value = false
  } catch (error) {
    console.error('表单验证失败:', error)
  }
}

const resetForm = () => {
  form.id = null
  form.name = ''
  form.serverAddr = ''
  form.clientId = 'client-' + Date.now()
}

onMounted(() => {
  // 加载连接列表
  loadConnections()
})

const loadConnections = () => {
  // TODO: 从后端加载连接列表
  connections.value = [
    {
      id: 1,
      name: '默认连接',
      serverAddr: '127.0.0.1:10001',
      clientId: 'client-001',
      status: 'disconnected',
      createTime: '2024-01-01 10:00:00'
    }
  ]
}
</script>

<style scoped>
.connections {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
