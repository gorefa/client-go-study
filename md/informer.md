> Kubernetes的其他组件都是通过client-go的Informer机制与Kubernetes API Server进行通信的。

### Informer核心组件

#### Reflector

Reflector用于监控（Watch）指定的Kubernetes资源，当监控的资源发生变化时，触发相应的变更事件，例如Added（资源添加）事件、Updated（资源更新）事件、Deleted（资源删除）事件，并将其资源对象存放到本地缓存DeltaFIFO中。

#### DeltaFIFO

DeltaFIFO可以分开理解，FIFO是一个先进先出的队列，它拥有队列操作的基本方法，例如Add、Update、Delete、List、Pop、Close等，而Delta是一个`资源对象存储`，它可以保存资源对象的操作类型，例如Added（添加）操作类型、Updated（更新）操作类型、Deleted（删除）操作类型、Sync（同步）操作类型等。

#### Indexer

Indexer是client-go用来存储资源对象并自带索引功能的本地存储，Reflector从DeltaFIFO中将消费出来的资源对象存储至Indexer。Indexer与Etcd集群中的数据完全保持一致。client-go可以很方便地从本地存储中读取相应的资源对象数据，而无须每次从远程Etcd集群中读取，以减轻Kubernetes APIServer和Etcd集群的压力。


sharedInformerFactory 实现了SharedInformerFactory interface的所有方法。
