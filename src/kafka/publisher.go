/*
 *  Copyright (c) 2011 NeuStar, Inc.
 *  All rights reserved.  
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at 
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *  
 *  NeuStar, the Neustar logo and related names and logos are registered
 *  trademarks, service marks or tradenames of NeuStar, Inc. All other 
 *  product names, company names, marks, logos and symbols may be trademarks
 *  of their respective owners.
 */

package kafka

type BrokerPublisher struct {
  broker *Broker
}

func NewBrokerPublisher(hostname string, topic string, partition int) *BrokerPublisher {
  return &BrokerPublisher{broker: newBroker(hostname, topic, partition)}
}

func (b *BrokerPublisher) Publish(message *Message) (int, error) {
  return b.BatchPublish(message)
}

func (b *BrokerPublisher) BatchPublish(messages ...*Message) (int, error) {
  conn, err := b.broker.connect()
  if err != nil {
    return -1, err
  }
  defer conn.Close()
  // TODO: MULTIPRODUCE
  request := b.broker.EncodePublishRequest(messages...)
  num, err := conn.Write(request)
  if err != nil {
    return -1, err
  }

  return num, err
}
