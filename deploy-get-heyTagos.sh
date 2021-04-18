gcloud functions deploy getHeyTagos \
  --entry-point GetTagosHttpServer \
  --runtime go113 \
  --trigger-http \
  --allow-unauthenticated \
  --memory 128MB \
  --region asia-northeast2 \
  --set-env-vars ENV_NAME=production