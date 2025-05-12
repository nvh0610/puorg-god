package kafka_manager

import "crawler_go_facebook/pkg/configs"

func NewConsumerProfile() (*KafkaConsumer, error) {
	opts := &configs.KafkaConfig{}
	opts.LoadConsumerEnvs()

	return newKafkaConsumer(opts, opts.TopicProfiles(), opts.GroupTopic(), &ProfileSerializer{})
}

func NewConsumerPost() (*KafkaConsumer, error) {
	opts := &configs.KafkaConfig{}
	opts.LoadConsumerEnvs()

	return newKafkaConsumer(opts, opts.TopicPosts(), opts.GroupTopic(), &PostSerializer{})
}

func NewConsumerComment() (*KafkaConsumer, error) {
	opts := &configs.KafkaConfig{}
	opts.LoadConsumerEnvs()

	return newKafkaConsumer(opts, opts.TopicComments(), opts.GroupTopic(), &CommentSerializer{})
}

func NewConsumerUserLiked() (*KafkaConsumer, error) {
	opts := &configs.KafkaConfig{}
	opts.LoadConsumerEnvs()

	return newKafkaConsumer(opts, opts.TopicUserLiked(), opts.GroupTopic(), &UserLikedSerializer{})
}

func NewConsumerUrlToId() (*KafkaConsumer, error) {
	opts := &configs.KafkaConfig{}
	opts.LoadConsumerEnvs()

	return newKafkaConsumer(opts, opts.TopicUrlToId(), opts.GroupTopic(), &UrlToIdSerializer{})
}

func NewConsumerReactor() (*KafkaConsumer, error) {
	opts := &configs.KafkaConfig{}
	opts.LoadConsumerEnvs()

	return newKafkaConsumer(opts, opts.TopicReactor(), opts.GroupTopic(), &ReactorSharerSerializer{})
}

func NewConsumerSharer() (*KafkaConsumer, error) {
	opts := &configs.KafkaConfig{}
	opts.LoadConsumerEnvs()

	return newKafkaConsumer(opts, opts.TopicSharer(), opts.GroupTopic(), &ReactorSharerSerializer{})
}
