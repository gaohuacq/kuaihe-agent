# kuaihe_agent

    快喝走服务发现的方式的demo程序 目前仅仅copy了 商品聚合服务的一部分接口

  ```code
  api类型：
  rest api：客户端前端直调，针对用户授权，如用户登录。
  open api：后端不同中心之间调用，针对内部clientId授权。服务发现方式调用，如订单中心调用商品中心的接口。
  public api：一般是前端无需登录的场景，查询数据。无需授权。
  manage api：只能管理中心web前端调用。
  开放能力接口：与外部系统交接数据，集成场景使用。通过开放网关访问。针对外部clientId授权。


restapi/v3/productgather/business/product/search  ========>  /openapi/v1/productgather/business/product/search

restapi/v2/productgather/frontCategory/category/topList ========>  /publicapi/v1/productgather/frontCategory/category/topList

/restapi/v1/productgather/frontCategory/category/storeChildCategoryList =======> /publicapi/v1/productgather/frontCategory/category/storeChildCategoryList

restapi/v3/productgather/business/product/queryProducts ======>  /openapi/v1/productgather/business/product/queryProducts

restapi/v2/productgather/business/product/couponProducts =====>  /publicapi/v1/productgather/business/product/couponProducts

restapi/v2/productgather/hotSearch/hotSearchListData  =======>  /openapi/v1/operation/hotSearchcom/hotSearchListData

restapi/v2/productgather/hotSearch/hotSearchInnerListData   =======>  /openapi/v1/operation/hotSearchcom/hotSearchInnerListData

restapi/v2/productgather/hotSearch/hotSearchKeyListData   =======>  /openapi/v1/operation/hotSearchcom/hotSearchKeyListData

restapi/v2/productgather/business/product/searchProductHotwords =======> 

restapi/v2/productgather/promotion/product/storeSecKillList =======> 

restapi/v2/productgather/business/product/collageDetail  =======> 

restapi/v2/productgather/business/product/collageList  =======> 

restapi/v2/productgather/business/product/preSaleList =======>  /openapi/v1/productgather/business/product/preSaleList

restapi/v2/productgather/category/couponCategoryList =======>  /publicapi/v1/productgather/category/couponCategoryList

restapi/v3/productgather/business/product/detail =======>  /publicapi/v1/productgather/business/product/detail

restapi/v2/productgather/business/product/preSaleDetail =======> 
  ```