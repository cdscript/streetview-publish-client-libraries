/*
 * Copyright 2017, Google Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * A client to Street View Publish API.
 *
 * <p>The interfaces provided are listed below, along with usage samples.
 *
 * <p>============================== StreetViewPublishServiceClient ==============================
 *
 * <p>Service Description: Publishes and connects user-contributed photos on Street View.
 *
 * <p>Sample for StreetViewPublishServiceClient:
 *
 * <pre>
 * <code>
 * try (StreetViewPublishServiceClient streetViewPublishServiceClient = StreetViewPublishServiceClient.create()) {
 *   Photo photo = Photo.newBuilder().build();
 *   Photo response = streetViewPublishServiceClient.createPhoto(photo);
 * }
 * </code>
 * </pre>
 */
package com.google.streetview.publish.v1;
