window._config = {
    cognito: {
        userPoolId: 'ap-northeast-1_dXWSmmhoC', // e.g. us-east-2_uXboG5pAb
        region: 'ap-northeast-1', // e.g. us-east-2
        clientId: '4g8tnjbiq7b3qf9chclibeh8ss' //is this used anywhere?
        // Add indentity poolID
        
    },
    s3: {
        BUCKET_NAME: 'web-test-avatar',
        REGION: 'ap-northeast-1',
        IDENTITY_POOL_ID: 'ap-northeast-1:559a80d5-f29f-4dfd-9eb7-6bbce6fae5ac'
    },
    api:{
        users : 'https://oc1lgx4pq6.execute-api.ap-northeast-1.amazonaws.com/stagging/users'
    }
};

